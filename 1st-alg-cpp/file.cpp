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

    for (const std::pair<const std::string, const int> &point : amount)
        std::cout << "Word \"" << point.first << "\" repeats " << point.second
                  << (point.second == 1 ? " time" : " times") << std::endl;

    return EXIT_SUCCESS;
}
