package godebug

import (
	"context"
	"path/filepath"

	"github.com/jmigpin/editor/util/goutil"
)

func SetupGoMods(ctx context.Context, cmd *Cmd, files *Files, mainFilename string, tests bool) error {
	dir := filepath.Dir(mainFilename)
	if tests {
		dir = mainFilename
	}

	// no go.mod defined (probably small simple file)
	if len(files.modFilenames) == 0 {
		// create go.mod file at tmp
		dirAtTmp := cmd.tmpDirBasedFilename(dir)
		content := "module example.com/main\n"
		if err := goutil.GoModCreateContent(dirAtTmp, content); err != nil {
			return err
		}

		if err := setupGoMod(ctx, cmd, files, dir); err != nil {
			return err
		}
		return nil
	}

	// updating all found go.mods, only the main one will be used
	for filename := range files.modFilenames {
		// update go.mod
		dir2 := filepath.Dir(filename)
		if err := setupGoMod(ctx, cmd, files, dir2); err != nil {
			return err
		}
	}
	return nil
}

func setupGoMod(ctx context.Context, cmd *Cmd, files *Files, dir string) error {
	// add to go.mod the godebugconfig location
	dirAtTmp := cmd.tmpDirBasedFilename(dir)
	if err := setupGodebugGoMod(ctx, cmd, dirAtTmp); err != nil {
		return err
	}

	// read go.mod
	goMod, err := goutil.ReadGoMod(ctx, dirAtTmp)
	if err != nil {
		return err
	}

	// update existing "replaces" relative dirs
	for _, rep := range goMod.Replace {
		np := rep.New.Path
		if !filepath.IsAbs(np) {
			abs, err := filepath.Abs(filepath.Join(dir, np))
			if err != nil {
				return err
			}
			if err := goutil.GoModReplace(ctx, dirAtTmp, rep.Old.Path, abs); err != nil {
				return err
			}
		}

	}

	// update/add "replaces" for the other mod files (annotated pkgs)
	for filename2 := range files.modFilenames {
		dir2 := filepath.Dir(filename2)
		if dir2 == dir { // same dir (same go mod file)
			continue
		}
		dirAtTmp2 := cmd.tmpDirBasedFilename(dir2)

		// read go.mod
		goMod2, err := goutil.ReadGoMod(ctx, dirAtTmp2)
		if err != nil {
			return err
		}

		// if gomod depends on gomod2
		for _, req := range goMod.Require {
			if req.Path == goMod2.Module.Path {
				if err := goutil.GoModReplace(ctx, dirAtTmp, req.Path, dirAtTmp2); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func setupGodebugGoMod(ctx context.Context, cmd *Cmd, dir string) error {
	{
		// require godebugconfig
		path2 := GoDebugConfigPkgPath + "@v0.0.0"
		if err := goutil.GoModRequire(ctx, dir, path2); err != nil {
			return err
		}
	}
	{
		// replace godebugconfig (point to tmp dir)
		oldPath := GoDebugConfigPkgPath
		newPath := filepath.Join(cmd.tmpDir, GoDebugConfigPkgPath)
		if err := goutil.GoModReplace(ctx, dir, oldPath, newPath); err != nil {
			return err
		}
	}
	{
		// replace debug (point to tmp dir)
		oldPath := DebugPkgPath
		newPath := filepath.Join(cmd.tmpDir, DebugPkgPath)
		// create go.mod file at tmp
		dirAtTmp := cmd.tmpDirBasedFilename(DebugPkgPath)
		content := "module " + DebugPkgPath + "\n"
		if err := goutil.GoModCreateContent(dirAtTmp, content); err != nil {
			return err
		}
		// replace
		if err := goutil.GoModReplace(ctx, dir, oldPath, newPath); err != nil {
			return err
		}
	}
	return nil
}
