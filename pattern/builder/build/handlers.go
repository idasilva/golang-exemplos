package build


func HandlerBuilder(director *director)  {
     director.setBuilder(
     	NewNormalBuilder(),
     	NewIglooBUilder(),
     	)
}




