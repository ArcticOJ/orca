package orca

import (
	"context"
	"github.com/ArcticOJ/blizzard/v0/config"
	"github.com/ArcticOJ/blizzard/v0/logger"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Orca struct {
	*discordgo.Session
}

var Instance *Orca

func Init(ctx context.Context) {
	conf := config.Config.Discord
	if conf == nil || strings.TrimSpace(conf.Token) == "" {
		return
	}
	logger.Orca.Info().Msg("initializing orca")
	s, e := discordgo.New("Bot " + conf.Token)
	logger.Panic(e, "failed to create a Discord session")
	go func() {
		<-ctx.Done()
		s.Close()
	}()
	Instance = &Orca{s}
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.Orca.Info().Msg("orca is ready")
	})
	logger.Panic(s.Open(), "failed to open a Discord session")
}
