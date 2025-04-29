# pongGo

A classic Pong game built using the Go programming language and the Raylib library.

## Features

*   Classic two-player Pong gameplay.
*   Score tracking up to 5 points.
*   Sound effects for game events (start, hit, score, game over).
*   Simple fade-in/fade-out effects for text.
*   Game over screen with options to restart or quit.

## Dependencies

*   **Go:** Version 1.x or later.
*   **Raylib:** The C library itself needs to be installed on your system. Follow the instructions for your OS on the [Raylib Wiki](https://github.com/raysan5/raylib/wiki).
*   **raylib-go:** The Go bindings for Raylib.

## Installation

1.  **Install Go:** Follow the official Go installation guide: [https://golang.org/doc/install](https://golang.org/doc/install)
2.  **Install Raylib:** Install the Raylib C library suitable for your operating system. See the [Raylib Wiki](https://github.com/raysan5/raylib/wiki) for detailed instructions.
3.  **Clone the repository:**
    ```bash
    git clone https://github.com/cw222kq/pongGo.git
    cd pongGo
    ```
4.  **Download Go dependencies:**
    ```bash
    go mod tidy
    ```
    This will download the `github.com/gen2brain/raylib-go/raylib` package.

## How to Play

1.  **Run the game:**
    ```bash
    go run .
    ```
2.  **Controls:**
    *   **Left Paddle:** `W` (Up), `S` (Down)
    *   **Right Paddle:** `Up Arrow` (Up), `Down Arrow` (Down)
3.  **Objective:** Be the first player to score 5 points. A point is scored when the opponent fails to return the ball.
4.  **Game Over:** When a player reaches 5 points, the game ends.
    *   Press `N` to start a New Game.
    *   Press `Q` to Quit the game.

## Code Structure

*   `main.go`: Entry point, game loop, window initialization, main event handling.
*   `game.go`: Core game logic, manages game state, ball, paddles, score, and transitions.
*   `ball.go`: Ball object logic, including movement and collision detection.
*   `paddle.go`: Paddle object logic, including movement.
*   `sound.go`: Handles loading and playing sound effects.
*   `assets/`: Contains sound files used by the game.

---