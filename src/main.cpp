#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <filesystem>

struct project_statistics {
  uintmax_t total_size = 0;
  size_t total_lines = 0;
  std::map<std::string, size_t> language_count;
};

size_t count_lines(const std::filesystem::path& file_path) {
  std::ifstream file(file_path);
  if (!file.is_open()) return 0;

  size_t line_count = 0;
  std::string line;
  while (std::getline(file, line)) {
    line_count++;
  }

  return line_count;
}

project_statistics analyze_project(const std::filesystem::path& project_path) {
  project_statistics stats;

  if (!std::filesystem::exists(project_path) || !std::filesystem::is_directory(project_path)) {
    std::cerr << "Invalid directory path: " << project_path << std::endl;
    return stats;
  }

  for (const auto& entry : std::filesystem::recursive_directory_iterator(project_path)) {
    if (std::filesystem::is_regular_file(entry.path())) {
      stats.total_size += std::filesystem::file_size(entry.path());

      size_t lines = count_lines(entry.path());
      stats.total_lines += lines;

      std::string ext = entry.path().extension().string();
      if (!ext.empty()) {
        stats.language_count[ext] += lines;
      }
    }
  }

  return stats;
}

int main(int argc, char* argv[]) {
  if (argc < 2) {
    std::cerr << "Error: Please provide a project path." << std::endl;
    return 1;
  }

  std::filesystem::path project_path(argv[1]);
  project_statistics resuls = analyze_project(project_path);

  std::cout << "\n====== PROJECT STATISTICS ======\n";
  std::cout << "Total Size: " << (resuls.total_size / 1024.0) << " KB\n";
  std::cout << "Lines of Code: " << resuls.total_lines << "\n";
  std::cout << "\nLanguage & File Makeup\n";

  for (const auto& [ext, count] : resuls.language_count) {
    std::cout << "- " << ext << ": " << (count / resuls.total_lines) * 100 << "%\n";
  }

  std::cout << "=================================\n";

  return 0;
}
