#!/usr/bin/env bash

tmux new-session -d "cd judger && cargo run" && \
  tmux split-window -h "cd server && go run main.go" && \
  tmux -2 attach-session -d

