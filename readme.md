## Welcome to Alchemy.

#### A Game Boy Emulator Written In Go.

Alchemy is a complete emulator of the 1989 Game Boy. It is written from scratch in Go. Currently, Alchemy is able to run all 32 KB roms such as Tetris, Dr. Mario, Asteroids and Tennis.


<img align="center" width="776" height="758" src="/demos/collage.gif">


### Installation

In order to run Alchemy, you will need to clone the repository and then compile the source code. You can do so by running the following commands. 

    git clone https://github.com/shafinsiddique/alchemy
    cd alchemy/src/alchemy
    go build -o alchemy
    ./alchemy

### Controls
    - Left Arrow
    - Right Arrow
    - Up Arrow
    - Down Arrow
    - Space for 'Select'
    - Enter for 'Start'
    - A for 'a'
    - S for 'b   
     
### Testing

Currently, Alchemy passes all of [Blargg's CPU Instruction Tests](https://github.com/retrio/gb-test-roms).

The PPU passes all parts of the [DMG-Acid2 Test](https://github.com/mattcurrie/dmg-acid2) except for the Window section due to the window layer not being implemented yet.

### Acknowledgements

This emulator could not have been completed without the support and 24/7 help from the amazing people on the [Emulation Development Discord Server](https://discord.gg/eZaeaxtQ). If you're building an emulator, i highly suggest you check this community out.

