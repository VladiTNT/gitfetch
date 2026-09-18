#include <fstream>
#include <iostream>
#include <filesystem>
#include <ostream>
#include <sstream>
#include <string>
#include <algorithm>

int get_num_lines(std::filesystem::path path) {
  std::stringstream stream;
  std::ifstream file(path);
  stream << file.rdbuf();
  file.close();
  std::string str = stream.str();
  return std::count(str.begin(), str.end(), '\n');
}

int main(int argc, char* argv[]) {
  if (argc < 2) {
    std::cout << "Too little arguments!" << std::endl;
    return 1;
  }

  for (const auto& entry : std::filesystem::directory_iterator(argv[1])) {
    if (entry.is_directory()) {
      continue;
    }

    std::cout << entry.path() << " Lines: " << get_num_lines(entry.path()) << std::endl;
  }

  return 0;
}
