#!/bin/bash

short_hash=$(git rev-parse --short HEAD)
docker build . -t "gcr.io/$PROJECT_ID/$SERVICE_NAME:$short_hash"

