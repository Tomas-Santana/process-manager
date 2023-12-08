package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)


type Process struct {
	Id 			string
	InitialTime int
	Time        int
	FinalTime   int
	RealTime	int
	Wait 		int
	WaitRatio	float64
	Done		bool
}

func main() {
	var processes []Process


	color.Red("Welcome to the Process Scheduler Simulator")
	defaultProcesses := []Process{
		newProcess("A", 2, 17),
		newProcess("B", 9, 47),
		newProcess("C", 8, 14),
		newProcess("D", 7, 32),
		newProcess("E", 6, 48),
		newProcess("F", 5, 23),
		newProcess("G", 40,13),
		newProcess("H", 4, 37),
		newProcess("I", 39, 24),
		newProcess("J", 38, 4),
		newProcess("K", 37, 25),
		newProcess("L", 36, 34),
		newProcess("M", 35, 26),
		newProcess("N", 34, 38),
		newProcess("O", 33, 15),
		newProcess("P", 32, 31),
		newProcess("Q", 31, 42),
		newProcess("R", 30, 21),
		newProcess("S", 3, 45),
		newProcess("T", 29, 43),
		newProcess("U", 28, 36),
		newProcess("V", 27, 22),
		newProcess("W", 26, 49),
		newProcess("X", 25, 18),
		newProcess("Y", 24, 39),
		newProcess("Z", 23, 27),
		newProcess("A1", 22, 46),
		newProcess("B1", 21, 16),
		newProcess("C1", 20, 30),
		newProcess("D1", 19, 44),
		newProcess("E1", 18, 35),
		newProcess("F1", 17, 20),
		newProcess("G1", 16, 50),
		newProcess("H1", 15, 19),
		newProcess("I1", 14, 33),
		newProcess("J1", 13, 41),
		newProcess("K1", 12, 28),
		newProcess("L1", 11, 40),
		newProcess("M1", 10, 29),
	}

	// defaultProcesses := []Process{
	// 	newProcess("A", 2, 1),
	// 	newProcess("B", 6, 6),
	// 	newProcess("C", 5, 3),
	// 	newProcess("D", 9, 5),
	// 	newProcess("E", 11, 4),
	// 	newProcess("F", 15, 10),
	// 	newProcess("G", 7, 9),
	// 	newProcess("H", 3, 7),
	// 	newProcess("I", 8, 2),
	// 	newProcess("J", 4, 8),
	// 	newProcess("K", 12, 11),
	// 	newProcess("L", 16, 12),
	// }

	fmt.Println("Input processes")
	for {
		var (
			id string
			initialTime int
			time int
		)
		color.Blue("Id (type 'end' if there are not more processes to add, type 'default' for default processes): ")
		fmt.Scan(&id)
		if id == "end" {
			break
		}
		if id == "default" {
			processes = defaultProcesses
			break
		}
		color.Green("Initial Time: ")
		fmt.Scan(&initialTime)
		color.Magenta("Time: ")
		fmt.Scan(&time)
		processes = append(processes, newProcess(id, initialTime, time))
	}

	if len(processes) == 0 {
		fmt.Println("No processes to schedule")
		return
	}
	fmt.Println("Choose a scheduling algorithm: ")
	color.Green("1. FIFO")
	color.Blue("2. LIFO")
	color.Red("3. Round Robin")
	color.Magenta("4. Compare")
	var option int
	fmt.Print("Option: ")
	fmt.Scan(&option)

	switch option {	
	case 1:
		start := time.Now()
		fifoManager(processes)
		elapsed := time.Since(start)
		fmt.Printf("Time elapsed: %s\n", elapsed)
	case 2:
		start := time.Now()
		lifoManager(processes)
		elapsed := time.Since(start)
		fmt.Printf("Time elapsed %s\n", elapsed)
	case 3:
		var quantum int
		fmt.Print("Quantum: ")
		fmt.Scan(&quantum)
		start := time.Now()
		rrManager(processes, quantum)
		elapsed := time.Since(start)
		fmt.Printf("Time elapsed %s\n", elapsed)
	case 4:
		var quantum int
		fmt.Print("Quantum: ")
		fmt.Scan(&quantum)
		compare(processes, quantum)
	default:
		fmt.Println("Invalid option")
	}
}

func fifoManager(processes []Process) (float64, float64, float64) {
	currentTime := 0
	done := 0

	for done < len(processes) {
		idle := true
		for i := range processes {
			process := &processes[i]
			if process.InitialTime <= currentTime && !process.Done {
				currentTime += process.Time
				process.Done = true
				calcProcess(process, currentTime)
				done += 1
				idle = false
				break
			}
		}
		if idle {
			currentTime += 1
		}
	}
	var (
		avgRealTime float64
		avgWaitRatio float64
		avgWait float64
	)
	avgRealTime, avgWaitRatio, avgWait = printProcesses(processes)

	return avgRealTime, avgWaitRatio, avgWait

}

func lifoManager(processes []Process) (float64, float64, float64) {
	idle := true
	currentTime := 0
	done := 0

	for done < len(processes) {
		for i := len(processes) - 1; i >= 0; i-- {
			process := &processes[i]
			if process.InitialTime <= currentTime && !process.Done {
				currentTime += process.Time
				process.Done = true
				calcProcess(process, currentTime)
				done += 1
				idle = false
				break
			}
		}
		if idle {
			currentTime += 1
		}
	}
	var (
		avgRealTime float64
		avgWaitRatio float64
		avgWait float64
	)
	avgRealTime, avgWaitRatio, avgWait = printProcesses(processes)

	return avgRealTime, avgWaitRatio, avgWait
}

func rrManager(processes []Process, quantum int) (float64, float64, float64) {
	idle := true
	currentTime := 0
	done := 0

	for done < len(processes) {
		for i := range processes {
			process := &processes[i]
			if process.InitialTime <= currentTime && !process.Done {
				if process.Time > quantum {
					currentTime += quantum
					process.Time -= quantum
					idle = false
				} else {
					currentTime += process.Time
					process.Done = true
					calcProcess(process, currentTime)
					done += 1
					idle = false
				}
			}
		}
		if idle {
			currentTime += 1
		}
	}
	var (
		avgRealTime float64
		avgWaitRatio float64
		avgWait float64
	)
	avgRealTime, avgWaitRatio, avgWait = printProcesses(processes)

	return avgRealTime, avgWaitRatio, avgWait


}

func calcProcess(process *Process, currentTime int) {
	process.FinalTime = currentTime
	process.RealTime = process.FinalTime - process.InitialTime
	process.Wait = process.RealTime - process.Time
	process.WaitRatio = float64(process.Time) / float64(process.RealTime)
}

func printProcesses(processes []Process) (float64, float64, float64){
	color.Red("|Id\t|Init\t|Time\t|Final\t|Real\t|Wait\t|WaitRatio\t|")
	var (
		avgRealTime float64
		avgWaitRatio float64
		avgWait float64
	)
	for _, p := range(processes) {
		fmt.Printf(
			"|%5s\t|%5d\t|%5d\t|%5d\t|%5d\t|%5d\t|%7.2f\t|\n",
			p.Id,
			p.InitialTime,
			p.Time,
			p.FinalTime,
			p.RealTime,
			p.Wait,
			p.WaitRatio,
		)
		avgRealTime += float64(p.RealTime)
		avgWaitRatio += p.WaitRatio
		avgWait += float64(p.Wait)
	}
	color.Blue("Average Real Time: %.5f\n", avgRealTime / float64(len(processes)))
	color.Green("Average Wait Ratio: %.5f\n", avgWaitRatio / float64(len(processes)))
	color.Yellow("Average Wait: %.5f\n", avgWait / float64(len(processes)))

	return avgRealTime, avgWaitRatio, avgWait
}

func compare(processes []Process, quantum int) {
	var (
		fifoProcesses []Process
		lifoProcesses []Process
		rrProcesses []Process
		fifoRT float64
		fifoW float64
		fifoWR float64
		lifoRT float64
		lifoW float64
		lifoWR float64
		rrRT float64
		rrW float64
		rrWR float64
	)
	for _, p := range(processes) {
		fifoProcesses = append(fifoProcesses, p)
		lifoProcesses = append(lifoProcesses, p)
		rrProcesses = append(rrProcesses, p)
	}

	fmt.Println("\nFIFO")
	fifoStart := time.Now()
	fifoRT, fifoWR, fifoW = fifoManager(fifoProcesses)
	fifoElapsed := time.Since(fifoStart)

	fmt.Println("\nLIFO")
	lifoStart := time.Now()
	lifoRT, lifoWR, lifoW = lifoManager(lifoProcesses)
	lifoElapsed := time.Since(lifoStart)

	fmt.Println("\nRR")
	rrStart := time.Now()
	rrRT, rrWR, rrWR = rrManager(rrProcesses, quantum)
	rrElapsed := time.Since(rrStart)

	// map to store the results

	var m map[string][]float64

	m = make(map[string][]float64)

	m["FIFO"] = []float64{fifoRT, fifoWR, fifoW}
	m["LIFO"] = []float64{lifoRT, lifoWR, lifoW}
	m["RR"] = []float64{rrRT, rrWR, rrW}


	fmt.Printf("FIFO Time elapsed: %s\n", fifoElapsed)
	fmt.Printf("LIFO Time elapsed: %s\n", lifoElapsed)
	fmt.Printf("RR Time elapsed: %s\n", rrElapsed)



}

func newProcess(id string, initialTime int, time int) Process {
	return Process{id, initialTime, time, 0, 0, 0, 0, false}
}

