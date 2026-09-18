# Compiler & Flags
CC := g++
FLAGS := -Wall -Wextra -O3

# Target Binary
TARGET := bin/gitfetch

# Directories
SRC_DIR := src
OBJS_DIR := bin/objs
BIN_DIR := bin

# Source files
SRCS := $(wildcard $(SRC_DIR)/*.cpp)

# Map cpp files to object files
OBJS := $(patsubst $(SRC_DIR)/%.cpp, $(OBJS_DIR)/%.o, $(SRCS))

# Base rule
all: $(TARGET)

# Link the final binary
$(TARGET): $(OBJS)
	@mkdir -p $(BIN_DIR)
	$(CC) $(FLAGS) $^ -o $@
	@echo "Built target: $@"

# Compile source files to object files
$(OBJS_DIR)/%.o: $(SRC_DIR)/%.cpp
	@mkdir -p $(OBJS_DIR)
	$(CC) $(FLAGS) -c $< -o $@

# Clean binary and build artifacts
clean:
	rm -rf $(OBJS_DIR) $(BIN_DIR)

# Phony declaration
.PHONY: all clean
