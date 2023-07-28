package main

import rl "github.com/gen2brain/raylib-go/raylib"
import "strconv"
import "fmt"
//ball values
type Ball struct{
	cordX int32
	cordY int32
	speedX int32
	speedY int32


}
func(p *Ball)  Draw(cordX int32,cordY int32){

	rl.DrawCircle(cordX,cordY,10.1,rl.White)

}
//paddle values
type Paddle struct{
	cordX int32
	cordY int32
	width int32
	height int32
	speed int32
}


func(p *Paddle) Draw(cordX int32,cordY int32,width int32,height int32){
	

	rl.DrawRectangle(cordX,cordY,width,height,rl.White)
}

// game function
func pongGame(){
	rl.InitWindow(800, 450, "PingPong-2P-game")//initialise window

	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	//variables
	width := int32(rl.GetScreenWidth())
	height := int32(rl.GetScreenHeight())
	var paddleLeftScore int = 0
	var paddleRightScore int = 0

	//ball variables
	ball := Ball{}	
	ball.cordX = 200
	ball.cordY = 100
	ball.speedX = 1
	ball.speedY = 1

	//paddleLeft variables
	paddleLeft := Paddle{}
	paddleLeft.cordX = 2
	paddleLeft.cordY = height /2
	paddleLeft.height = 100
	paddleLeft.width = 10
	paddleLeft.speed = 15

	//paddleRight variables
	paddleRight := Paddle{}
	paddleRight.cordX = width -12
	paddleRight.cordY = height / 2
	paddleRight.height = 100
	paddleRight.width = 10
	paddleRight.speed = 15

	//end of variables



	//creating window
	for !rl.WindowShouldClose() {
		//game Physics
		for i := 0 ; i < 10; i++{
			ball.cordX += ball.speedX
			ball.cordY += ball.speedY
		}
		if ball.cordY == height{
			ball.cordY = height
			ball.speedY *= -1
		}
		if ball.cordY == 0 {
			ball.cordY = 0
			ball.speedY *= -1
		}
		//check for collision on right paddle
		if rl.CheckCollisionCircleRec(rl.Vector2{float32(ball.cordX),float32(ball.cordY)},10.1, rl.Rectangle{float32(paddleRight.cordX),float32(paddleRight.cordY),float32(paddleRight.width),float32(paddleRight.height)}){
	
		ball.speedX *= -1
		paddleRightScore += 1
		}

		//check for collision on left paddle
		if rl.CheckCollisionCircleRec(rl.Vector2{float32(ball.cordX),float32(ball.cordY)},10.1, rl.Rectangle{float32(paddleLeft.cordX),float32(paddleLeft.cordY),float32(paddleLeft.width),float32(paddleLeft.height)}){
			ball.speedX *= -1
			paddleLeftScore += 1
			}

		//end of physics
	//controles
	if  rl.IsKeyDown(87){
		//the W key
		if paddleLeft.cordY <= 0 {
			paddleLeft.cordY = 0
			fmt.Println("reachedup")
		}else{
			paddleLeft.cordY -= paddleLeft.speed
		}
	}
	if rl.IsKeyDown(83){	
		if paddleLeft.cordY > height - paddleLeft.height{
			paddleLeft.cordY = height - paddleLeft.height
			fmt.Println("reached")
		}else{
			paddleLeft.cordY += paddleLeft.speed
		}
	}


	if  rl.IsKeyDown(265){
		//the Up key
		if paddleRight.cordY <= 0{
			paddleRight.cordY = 0
		}else{
			paddleRight.cordY -= paddleRight.speed
		}
	}
	if rl.IsKeyDown(264){
		//the Down key	
		if (paddleRight.cordY > height - paddleRight.height){
			paddleRight.cordY = height - paddleRight.height
		}else{
			paddleRight.cordY += paddleRight.speed
		}
	}
	if rl.IsKeyDown(81){
		//close window
		rl.CloseWindow()
	}

	//end of controles	
	
	ga := int(rl.GetFrameTime())
	age := strconv.Itoa(ga)
	
		//begin Rendering on screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		paddleRight.Draw(paddleRight.cordX,paddleRight.cordY,paddleRight.width,paddleRight.height)
		ball.Draw(ball.cordX,ball.cordY)
		paddleLeft.Draw(paddleLeft.cordX,paddleLeft.cordY,paddleLeft.width,paddleLeft.height)

		rl.DrawText(age,width /2,4,50,rl.White)

		rl.DrawText(strconv.Itoa(paddleLeftScore),4,4,30,rl.White)
		rl.DrawText(strconv.Itoa(paddleRightScore),width -20,4,30,rl.White)

		if ball.cordX > width || ball.cordX < 0 {
			rl.DrawText("Game Over",width /2,height /2,50,rl.Red)
		}
		rl.EndDrawing()
		//end Rendering
	
	}
	//end of window

}
//main function
func main() {
	pongGame()

}
