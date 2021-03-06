#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <algorithm>

std::vector<std::string> words;
std::vector<int> indexes;

int main()
{
    std::ifstream file("words.txt");
    std::string str;
    while (std::getline(file, str))
        words.push_back(str);

    std::sort(words.begin(), words.end());

    for (int i = 0; i < words.size() - 1; i++)
        if (words.at(i) != words.at(i + 1))
            indexes.push_back(i + 1);

    indexes.insert(indexes.begin(), 0);
    indexes.push_back(words.size());

    for (int i = 0; i < indexes.size() - 1; i++)
        std::cout << "Word " << words.at(indexes.at(i))
                  << " repeats " << indexes.at(i + 1) - indexes.at(i)
                  << (indexes.at(i + 1) - indexes.at(i) == 1 ? " time" : " times")
                  << "\n";

    return EXIT_SUCCESS;
}
