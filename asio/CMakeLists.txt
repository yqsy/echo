project(echo)
cmake_minimum_required(VERSION 2.8)

include(conanbuildinfo.cmake)
CONAN_BASIC_SETUP()

add_executable(client client.cpp)
target_link_libraries(client ${CONAN_LIBS} pthread)

add_executable(server server.cpp)
target_link_libraries(server ${CONAN_LIBS} pthread)
