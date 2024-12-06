package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	//"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	treeItemView = tview.NewTextView()
)

func PickFileView(nextSlide func()) (title string, info string, content tview.Primitive) {

	rootDir := "."
	root := tview.NewTreeNode(rootDir)
	root.SetColor(tcell.ColorRed.TrueColor())
	tree := tview.NewTreeView()
	tree.SetRoot(root)
	tree.SetCurrentNode(root)

	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *tview.TreeNode, path string) {
		pathInfo, err := os.Stat(path)
		if err != nil {
			log.Print(err)
			return
		}
		if !pathInfo.IsDir() {
			return
		}
		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {

			isDir := file.IsDir()
			isCSV := strings.HasSuffix(file.Name(), ".csv")

			if !isDir && !isCSV {
				continue
			}

			node := tview.NewTreeNode(file.Name())
			// filter
			node.SetReference(filepath.Join(path, file.Name()))
			node.SetSelectable(true)
			if isDir {
				node.SetColor(tcell.ColorGreen.TrueColor())
			} else {
				node.SetColor(tcell.ColorPurple.TrueColor())
			}
			target.AddChild(node)
		}
	}

	// Add the current directory to the root node.
	add(root, rootDir)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		path := reference.(string)
		if strings.HasSuffix(path, ".csv") {
			treeItemView.SetText("[purple]" + path)
			return
		}

		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	treeItemView.SetWrap(false)
	treeItemView.SetDynamicColors(true)
	treeItemView.SetText("[red]selectedItem")
	treeItemView.SetBorderPadding(1, 1, 2, 0)

	btn := tview.NewButton("Load")

	btn.SetSelectedFunc(func() {
		txt := treeItemView.GetText(false)
		treeItemView.SetText(txt + " loadding...")
	})

	flex := tview.NewFlex()
	flex.AddItem(tree, 0, 1, true)
	flex.AddItem(treeItemView, codeWidth, 1, false)
	flex.AddItem(btn, 10, 1, false)

	return "FileImport", "", flex
}
