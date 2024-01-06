// NOTE: 62 CHARS ASCII

package main

import (
	"fmt"

	"github.com/liamg/gobless"
)

type Location struct {
	Name   string
	StartX int
	EndX   int
	StartY int
	EndY   int
}

const (
	maxBoardX = 10
	minBoardX = -10
	maxBoardY = 10
	minBoardY = -10
)

var (
	gui           *gobless.GUI
	controlBox    *gobless.TextBox
	controlText   string
	gameText      string
	userPositionX int
	userPositionY int
	locations     []Location
)

func main() {
	startGame()
}

func startGame() {
	gui = gobless.NewGUI()
	if err := gui.Init(); err != nil {
		panic(err)
	}
	defer gui.Close()

	controlText = ""
	gameText = ""
	userPositionX = 0
	userPositionY = 0

	locations = []Location{
		{
			Name:   "Dungeon",
			StartX: 3,
			EndX:   10,
			StartY: 3,
			EndY:   -3,
		},
		{
			Name:   "Town",
			StartX: 0,
			EndX:   10,
			StartY: 10,
			EndY:   5,
		},
	}

	closeGame()
	moveNorth()
	moveSouth()
	moveWest()
	moveEast()
	refreshGUI()

	gui.HandleResize(func(event gobless.ResizeEvent) {
		refreshGUI()
	})

	gui.Loop()
}

func refreshGUI() {
	rows := []gobless.Component{
		gobless.NewRow(
			gobless.GridSizeThreeQuarters,
			gobless.NewColumn(
				gobless.GridSizeTwoThirds,
				renderGameBox(gameText),
			),
			gobless.NewColumn(
				gobless.GridSizeOneThird,
				gobless.NewRow(
					gobless.GridSizeFull,
					gobless.NewColumn(
						gobless.GridSizeFull,
						renderInformationBox(),
						renderPositionBox(userPositionX, userPositionY),
					),
				),
			),
		),
		gobless.NewRow(
			gobless.GridSizeOneQuarter,
			gobless.NewColumn(
				gobless.GridSizeOneSixth,
				renderControlBox(controlText),
			),
		),
	}

	gui.Render(rows...)
}

func moveNorth() {
	gui.HandleKeyPress(gobless.KeyUp, func(event gobless.KeyPressEvent) {
		controlText = `
		 _  _    ___    ___   _____   _  _ 
		| \| |  / _ \  | _ \ |_   _| | || |
		| .  | | (_) | |   /   | |   | __ |
		|_|\_|  \___/  |_|_\   |_|   |_||_|
	   `
		gameText = `
		 ..............................................................
		 ..............................................................
		 .............................=*++**+..........................
		 .............................@#+=+#%:.........................
		 ........................:-=#%%#***%%%#*-:.....................
		 .........................:+*#%%%%%%%#*+:......................
		 .............................+@@@@@+:.........................
		 ............................:-%####*:.........................
		 ..........................-===-:-**##*-.......................
		 ........................:*+-::-+**+*::+-......................
		 ......................:==**==**#+=*@+--+=.....................
		 ......................=++%#**%%+**#@#+-=+:....................
		 ....................:+==#@%###%***%@@%*-=+=...................
		 ...................:==*#%#%#%%#*##*%@=##*-+*-.................
		 ...................-+**%*-%@##**###@@:=%##+==.................
		 ...................:=+#=-#%%*##*%#%@%:.:=#=---................
		 ..................-++*-.+@#*+++=+#%##*==:.=+==-...............
		 ..................+%%#-.+#*++%*--=+#@@%*+..+@#*-..............
		 .................-+#@@*=###%%@@%##**@@#%#..+@%*=:.............
		 .................=+**+#%%#*+*@@%#@%#@@#%+..:+**+-.............
		 ...................::=#%#+####%..+%@#%@@#=..:+*-..............
		 ......................**=-=*#@:....%@@@%#+....................
		 ....................:***=+*%#:.....-*%%%#*=:..................
		 ..................:=++++*#%%=.......:%@%%#*+:.................
		 .................-+=-:-**#%=.........+@@###**:................
		 ................=%%%@%++%@*...........+@%%#+==:...............
		 ................=*=+#@%#*:...........:*%%#*+=+=...............
		 ..............:=++*%@@*-..............:%@#**#%=.-++:..........
		 .............-#%##%%@=..................:=%@@%##*++#-.........
		 .............:*@@@%%#-....................=@@@%%%%%#:.........
		 .......:.......+@@@%%%*:...................+@%%%%%=..-........
		 ...--+*#%#==+*+*@@@@@%@+=---------=--------+@@@@*--+#=*%#*++-.
		 ...:=-====+++====================++=====++=================--:
		 .......:...:......::.:..:::.:..:::..:...::........::.:...:....
	`
		if userPositionY < maxBoardY {
			userPositionY++
		}
		checkUsersLocation()
		refreshGUI()
	})
}

func moveSouth() {
	gui.HandleKeyPress(gobless.KeyDown, func(event gobless.KeyPressEvent) {
		controlText = `
		 ___    ___    _   _   _____   _  _ 
		/ __|  / _ \  | | | | |_   _| | || |
		\__ \ | (_) | | |_| |   | |   | __ |
		|___/  \___/   \___/    |_|   |_||_|
	   `
		gameText = `
	 	..............................................................
	 	..............................................................
	 	.............................:................................
	 	...........................-=--:-++...........................
	 	...........................+-===++*...........................
	 	........................-===--====+++=:.......................
	 	.......................-+##**====+*##*-.......................
	 	...........................-*::::++...........................
	 	............................=****+............................
	 	.........................:=+++=-*+++-:........................
	 	........................=:=+==+++-++:=:.......................
	 	.......................:-:+*-.=:..**-:=.......................
	 	.......................=::**:.*...+*=:+.......................
	 	......................:=-=*+..:..:=#+---......................
	 	.....................:===+*+-::::-=#*==+-.....................
	 	......................==+*#+-:----=#**==:.....................
	 	.....................:=-**#***+*******=-=.....................
	 	.....................==+*+**++++++*#**+=+.....................
	 	.....................++**#*=-::::--+*-**+-....................
	 	.....................==++*=--+**+--+*-+=+:....................
	 	.....................+#=+*=-:=*++::+*=*#+:....................
	 	......................-=+*=-:+=:*-:=+:........................
	 	......................::-+-:-*:.*-:=+-........................
	 	........................:=-:-*..+=-=+=........................
	 	........................-=--++..-*==++........................
	 	........................=++++=...*=+++........................
	 	........................+=--*:...*=-++:.......................
	 	.......................:+=-=*....*===+-.......................
	 	.......................=+=++*....+*+=+*.......................
	 	.......................:++=#+....-#++*=.......................
	 	.......................-****=....:#***+.......................
	 	......................+++**#-.....#**++*:.....................
	 	.............::.::-:.:------------------::.::..::.............
	 	..............................................................
	  `
		if userPositionY > minBoardY {
			userPositionY--
		}
		checkUsersLocation()
		refreshGUI()
	})
}

func moveEast() {
	gui.HandleKeyPress(gobless.KeyRight, func(event gobless.KeyPressEvent) {
		controlText = `
		 ___     _     ___   _____ 
		| __|   /_\   / __| |_   _|
		| _|   / _ \  \__ \   | |  
		|___| /_/ \_\ |___/   |_|  
		`
		gameText = `
		 ..............................................................
		 ................................:--::.........................
		 ...............................+@#*=--:.......................
		 ............................:::%@@%##+=:......................
		 ..........................:#@@@@@@@@%#+=++=:..................
		 ...........................::-*@@%%+-=++-::...................
		 ...........................:-=+*#=#%+--.......................
		 .........................:*+--=*+=*+--:.......................
		 ......................:-=+=:::-=##%*-.......:.................
		 ....................:+***=++++*#=##*.......:=-:--:............
		 ...................:=##**=##%@%+-##*......::::=*=:............
		 ...................:+*#+:+@@****###+....:==-=*+:..............
		 .....................:*#*=-##**##%##=::-+##*+=:...............
		 .......................:-+#+-=#**@*+=++**++#*-................
		 .....................--==#%*==%##@##*+--=+==:.................
		 ....................=*##%@%%@#%##-::..........................
		 ..................:=**#@@@@%%%====::..........................
		 ..................:+#%@@@%#*+==-:::==--::.....................
		 ....................::-#@@@##**++=+=+--+-.....................
		 .....................=%%#%@@@#%#%%#%+*#::--:..................
		 ....................-#%##%%@*:.::*%###++*+=-..................
		 ....................:#@%#%%%:.....:--:=#+--++:................
		 .............:=**+++++=*%%*-..........:+#+--:--:..............
		 ............:+%+=+******+:............:=%#**+=+=..............
		 ..........:-+*##**#%*==-:..............-#@##%%*-:.....::......
		 .......:-=*##%@%#%#=:...................-**#@%*++=:::--+-:....
		 .....:=#*+#%%%=:.::........................:=#%#***#*+*%-.....
		 ......:+%%#%#-...............................:*%%**%%%#-......
		 ........-+%%%*-:..:......::......:...:........-#%%@*-::.......
		 .....:=::-#@%#%*:-=-::---==-.....-++=:--:-..::::*#-...-=+:....
		 ..:=-=*#*##**+**++*#**#*###**%##*++**##*#%%#**####%%#++====--:
		 ...:::-::-::-=------:::--:-=-:-:-----=:--::-:::-===-:---:::::.
		 ..............................................................
		 ..............................................................
		`
		if userPositionX < maxBoardX {
			userPositionX++
		}
		checkUsersLocation()
		refreshGUI()
	})
}

func moveWest() {
	gui.HandleKeyPress(gobless.KeyLeft, func(event gobless.KeyPressEvent) {
		controlText = `
		__      __  ___   ___   _____ 
		\ \    / / | __| / __| |_   _|
		 \ \/\/ /  | _|  \__ \   | |  
		  \_/\_/   |___| |___/   |_|  
		`
		gameText = `
		 ..............................................................
		 .........................::--:................................
		 .......................:--=+#%+...............................
		 ......................:=+*#%@@%:::............................
		 ..................:=++=+#%@@@@@@@@%:..........................
		 ...................::-++=-+%%@@*-::...........................
		 .......................--=##+#*++-:...........................
		 .......................::-+*=+*==-+*-.........................
		 .................:.......-+%##+-:::=+==:......................
		 ............:--:-=-.......*##=#*++++=***+:....................
		 ............:=*=::::......*##-=#@@##=**##=:...................
		 ..............:+*=-==:....=###****@@+:+#*+:...................
		 ...............:=+*##+-::=##%##**##-=*#*:.....................
		 ................-*#*+**++=+*%**#=-+#*-:.......................
		 .................:=++=--+*##@##%==*%#==--.....................
		 ..........................::-##@#@%%@%##*=....................
		 ..........................::====%%%@@@@#**=:..................
		 .....................::--==:::-==+*#%@@@%#+:..................
		 .....................-=--+=+=++**##@@@%-::....................
		 ..................:--::#*+##%%#%#@@@%#%%=.....................
		 ..................-=+*++###%#-:..+@%%##%#-....................
		 ................:+*--+#=:--:.....:%#%#%@%:....................
		 ..............:--:--+#+:..........:*%%*=+++++**=:.............
		 ..............=+=+**#%=:............:+#*****+=+%+:............
		 .......:.....:-*%%##%#-..............:-==*%#**##*+-:..........
		 ....:-+--:::-++*%@#**-...................:=#%#%@%##*=-:.......
		 .....-%*+*#***#%#=:........................::.:=#%%#+*#=:.....
		 ......-*%%%**#%*:...............................-*%#%%*:......
		 ........:-*@%%#-........:...:......::......:..:-+%%%+-:.......
		 ....:+=-...-#*::::..----:-++-.....:==---::--=:+%#%@#-::-:.....
		 :--====++*%%####**#%%#**#**++*##%**###*#*+#*++**+***#*#*=-=:..
		 .:::::---:-===-:::-::--:=-----:-:-=-:--:::------=-::-::-:::...
		 ..............................................................
		 ..............................................................
		 `

		if userPositionX > minBoardX {
			userPositionX--
		}
		checkUsersLocation()
		refreshGUI()
	})
}

func closeGame() {
	gui.HandleKeyPress(gobless.KeyCtrlC, func(event gobless.KeyPressEvent) {
		gui.Close()
	})
}

func renderGameBox(text string) *gobless.TextBox {
	gameBox := gobless.NewTextBox()
	gameBox.SetBorderColor(gobless.ColorGreen)
	gameBox.SetTitle("Game")
	gameBox.SetText(text)
	gameBox.SetTextWrap(true)
	return gameBox
}

func renderControlBox(text string) *gobless.TextBox {
	controlBox = gobless.NewTextBox()
	controlBox.SetText(text)
	controlBox.SetBorderColor(gobless.ColorDarkRed)
	return controlBox
}

func renderPositionBox(boardX int, boardY int) *gobless.TextBox {
	positionBox := gobless.NewTextBox()
	positionBox.SetTitle("Position")
	formattedText := fmt.Sprintf("X: %d, Y: %d", boardX, boardY)
	positionBox.SetText(formattedText)
	positionBox.SetBorderColor(gobless.Color100)
	return positionBox
}

func renderInformationBox() *gobless.TextBox {
	informationBox := gobless.NewTextBox()
	informationBox.SetTitle("Information")
	informationBox.SetText(`Objective: Find the hidden treasure.
    
    Controls:
        
    "Key ↑" to move North
    "Key ↓" to move South
    "Key ←" to move West
    "Key →" to move East
    
    Navigate wisely and watch for clues. Your adventure begins now! Good luck!`)
	return informationBox
}

func checkUsersLocation() {
	for _, loc := range locations {
		if userPositionX >= loc.StartX && userPositionX <= loc.EndX &&
			userPositionY >= loc.EndY && userPositionY <= loc.StartY {
			switch loc.Name {
			case "Dungeon":
				gameText = `
				88888888888888888888888888888888888888888888888888888888888888888888888
				88.._|      |  -.  |  .  -_-_ _-_  _-  _- -_ -  .'|   |.'|     |  _..88
				88    -.._  |    | !  | .  -_ -__ -_ _- _-_-  .'  |.;'   |   _.!-'|  88
				88      |  -!._  |   ;!  ;. _______________ ,'| .-' |   _!.i'     |  88
				88..__  |     | -!._ |  .| |_______________||."'|  _!.;'   |     _|..88
				88   |  "..__ |    | ";.| i|_|MMMMMMMMMMM|_|'| _!-|   |   _|..-|'    88
				88   |      |  --..|_ |  ;!|l|MMoMMMMoMMM|1|.'j   |_..!-'|     |     88
				88   |      |    |   | -,!_|_|MMMMP'YMMMM|_||.!-;'  |    |     |     88
				88___|______|____!.,.!,.!,!|d|MMMo * loMM|p|,!,.!.,.!..__|_____|_____88
				88      |     |    |  |  | |_|MMMMb,dMMMM|_|| |   |   |    |      |  88
				88      |     |    |..!-;'i|r|MPYMoMMMMoM|r| | -..|   |    |      |  88
				88      |    _!.-j'  | _!,"|_|M<>MMMMoMMM|_||!._|   i-!.._ |      |  88
				88     _!.-'|    | _."|  !;|1|MbdMMoMMMMM|l| .|  -._|    |  -.._  |  88
				88..-i'     |  _.''|  !-| !|_|MMMoMMMMoMM|_|.| -. |   ._ |     |  "..88
				88   |      |.|    |.|  !| |u|MoMMMMoMMMM|n|| . | !   |  ".    |     88
				88   |  _.-'  |  .'  |.' |/|_|MMMMoMMMMoM|_|! | !   ,.|    |-._|     88
				88  _!"'|     !.'|  .'| .'|[@]MMMMMMMMMMM[@] \|   . |  ._  |    -._  88
				88-'    |   .'   |.|  |/| /                 \| .  | !    |.|      | -88
				88      |_.'|   .' | .' |/                   \  \ |   .  |  ._-   |  88
				88     .'   | .'   |/|  /                     \ | !   | .|     .  |  88
				88  _.'     !'|   .' | /                       \|     |   .    | .|  88
				88888888888888888888888888888888888888888888888888888888888888888888888
				`
			case "Town":
				gameText = "Town"
			}
		}
	}
}
