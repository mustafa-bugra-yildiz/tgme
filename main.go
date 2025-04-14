package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type config struct {
	token    string
	receiver int64
}

func main() {
	cfg, err := readConfig();
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: reading config: %s\n", err.Error());
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: creating bot: %s\n", err.Error());
		os.Exit(1)
	}

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: tgme <message>\n")
		os.Exit(1)
	}

	message := tgbotapi.NewMessage(cfg.receiver, args[0])
	_, err = bot.Send(message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: sending message: %s\n", err.Error());
		os.Exit(1)
	}
}

func readConfig() (*config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(homeDir, ".config", "tgme", "config")

	contents, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	kv := make(map[string]string)
	lines := bytes.Split(contents, []byte("\n"))
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := bytes.Split(line, []byte("="))
		if len(parts) != 2 {
			return nil, errors.New("impartial configuration line")
		}

		key := bytes.TrimSpace(parts[0])
		value := bytes.TrimSpace(parts[1])
		kv[string(key)] = string(value)
	}

	token, ok := kv["token"]
	if !ok {
		return nil, errors.New("token missing")
	}

	receiverString, ok := kv["receiver"]
	if !ok {
		return nil, errors.New("receiver missing")
	}

	receiver, err := strconv.ParseInt(receiverString, 10, 64)
	if err != nil {
		return nil, errors.New("receiver must be int64")
	}

	return &config{
		token: token,
		receiver: receiver,
	}, nil
}
