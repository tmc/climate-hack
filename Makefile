.PHONY: init run lint test clean deps python-deps

# default command
all: init

init:
	make deps
	make python-deps

# install sys dependencies
deps:
	brew install poetry ruff

# Install Python dependencies
python-deps:
	poetry install

# run FastAPI server
run:
	poetry run uvicorn app.main:app --reload

# use Ruff for Python linting
lint:
	@poetry run ruff check .

# run tests with pytest (assuming you have tests set up)
test:
	@poetry run pytest

# remove generated files
clean:
	find . -name "__pycache__" -exec rm -r {} \+
	find . -name "*.pyc" -exec rm -f {} \+