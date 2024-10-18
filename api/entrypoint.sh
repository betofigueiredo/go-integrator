#!/bin/sh

alembic upgrade head

fastapi dev src/main.py --host 0.0.0.0 --port 8000
