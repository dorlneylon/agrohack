package points

import (
	`bytes`
	`context`
	`log`
	`log/slog`
	`os/exec`
	`time`

	`dorlneylon/agrofront/internal/config`
)

type Service struct {
	log *slog.Logger
	cfg config.Service
}

func New(log *slog.Logger, cfg config.Service) *Service {
	return &Service{
		log: log,
		cfg: cfg,
	}
}

func (s *Service) ProceedAllByDateTime(ctx context.Context, date time.Time) error {
	err := shellout(s.cfg.Command, s.cfg.File, date.Format("02_01_2006_15_04"))
	return err
}

const ShellToUse = "cmd"

func shellout(command string, arg ...string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	//log.Println(command, arg)
	cmd := exec.Command(command, arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	log.Println(stdout.String())
	//
	log.Println(stderr.String())
	return err
}
