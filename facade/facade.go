package facade

import "fmt"

func NewComputerFacade() *computerFacade {
	return &computerFacade{
		cpu:       &cpu{},
		memory:    &memory{},
		hardDrive: &hardDrive{},
	}
}

type computerFacade struct {
	cpu       *cpu
	memory    *memory
	hardDrive *hardDrive
}

func (cf *computerFacade) Start() {
	position := "0x00"
	cf.cpu.Freeze()
	cf.memory.Load(position, cf.hardDrive.read(100, 1024))
	cf.cpu.Jump(position)
	cf.cpu.Execute()
}

type cpu struct{}

func (c *cpu) Freeze() {
	fmt.Println("Freezing processor.")
}

func (c *cpu) Jump(position string) {
	fmt.Printf("Jumping to: %v.\n", position)
}

func (c *cpu) Execute() {
	fmt.Println("Executing.")
}

type memory struct{}

func (m *memory) Load(position, data string) {
	fmt.Printf("Loading from %v data: '%v'.\n", position, data)
}

type hardDrive struct{}

func (h *hardDrive) read(lba, size int) string {
	return fmt.Sprintf("Some data from sector %d with size %d", lba, size)
}
