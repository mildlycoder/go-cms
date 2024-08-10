# Define directories
PUBLIC_DIR = public
TAILWIND_INPUT = $(PUBLIC_DIR)/styles.css
TAILWIND_OUTPUT = $(PUBLIC_DIR)/output.css

# Define the tailwindcss command
TAILWIND_CMD = npx tailwindcss -i $(TAILWIND_INPUT) -o $(TAILWIND_OUTPUT) --watch

# Define the air command
AIR_CMD = air

# Install dependencies
.PHONY: install
install:
	npm install
	go mod tidy

# Run Tailwind CSS and air concurrently
.PHONY: dev
dev:
	$(TAILWIND_CMD) & $(AIR_CMD)

# Run only Tailwind CSS
.PHONY: tailwind
tailwind:
	$(TAILWIND_CMD)

# Run only the Go server with air
.PHONY: server
server:
	$(AIR_CMD)

# Clean up generated files
.PHONY: clean
clean:
	rm -f $(TAILWIND_OUTPUT)
	rm -rf tmp

# Build the Go application without running air
.PHONY: build
build:
	templ generate
	go build -o ./tmp/main .

# Rebuild the project (clean, then build)
.PHONY: rebuild
rebuild: clean build

