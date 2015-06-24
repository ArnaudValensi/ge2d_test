package main

import (
	"fmt"
	"binpix/ge2d"
)

func main () {
	fmt.Print("Hell !!\n")

	ge2d.Run()
	return

	// Init managers
	renderManager := ge2d.NewRenderManager()
	sceneManager := ge2d.NewSceneManager()

	// Create an object
	obj1 := ge2d.NewObject(1)
	// Create a render component
	spr := renderManager.CreateRenderComponent("walk_face")
	// Add it to the object
	obj1.AddComponent(spr)
	
	// Add object to the scene
	rootNode := sceneManager.GetRootNode()
	rootNode.AttachObject(obj1)

	// scene.AddObject(obj1)

	msg := ge2d.NewSetPositionMessage(1, ge2d.Vector2d {23, 69})
	sceneManager.HandleMessage(msg)

	// Free resources
	// renderManager.Quit()

	// var inode ge2d.INode = ge2d.NewBaseNode("root", nil, ge2d.Vector2d {5, 6})
	// fmt.Print("inode: ", inode)

	/////////////////////////////////////////////

	// scene := ge2d.NewSceneManager()
	// obj1 := ge2d.NewBaseNode("node", nil, ge2d.Vector2d {5, 6})
	// // obj1 := ge2d.NewObject(1)
	// msg := ge2d.NewSetPositionMessage(1, ge2d.Vector2d {23, 69})
	// scene.AddNode(obj1)
	// scene.HandleMessage(msg)

	// fmt.Print("inode: ", inode)

	/////////////////////////////////////////////

	// b := ge2d.NewBaseMessage(1, ge2d.MSG_SET_POSITION)
	// var i ge2d.IMessage = b
	

	// ge2d.Run()

	// sprite := ge2d.SpriteLoader {}

	// sprite := ge2d.NewSpriteLoader()
	// sprite.Load("./test.gspr")
}
