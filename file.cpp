#include <iostream>
#include <fstream>
#include <string>
#include <map>

std::map<std::string, int> amount;

int main()
{
    std::ifstream file("words.txt");
    std::string str;

    while (std::getline(file, str))
        if (amount.find(str) == amount.end())
            amount.insert(std::pair<std::string, int>(str, 1));
        else
            amount.at(str)++;

    for (auto &t : amount)
        std::cout << t.first << " - " << t.second << std::endl;

    return EXIT_SUCCESS;
}
