package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KlatterAB/klatter-burton/cfg"
	"github.com/KlatterAB/klatter-burton/util"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

type TimerParams struct {
	Duration      time.Duration
	SoundFilePath string
}

func StartTimer(params TimerParams) error {

	fmt.Printf("Starting a %.0f minutes long work session\n", params.Duration.Minutes())
	timer := time.NewTimer(params.Duration)

	<-timer.C
	err := playSound(params.SoundFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Time's up!")
	if cfg.Cfg.Notifications == true {
		util.Notify("Time to switch!", fmt.Sprintf("Your %.0f minutes long work session is over", params.Duration.Minutes()))
	}

	return nil
}

func playSound(soundFilePath string) error {
	f, err := os.Open(soundFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	s, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer s.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
		return err
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		done <- true
	})))
	<-done

	return nil
}
