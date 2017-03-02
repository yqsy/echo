conan install

BUILD_DIR="./build"
mkdir -p $BUILD_DIR

cmake -DCMAKE_BUILD_TYPE="release" -H"./" -B$BUILD_DIR

cmake --build $BUILD_DIR --target all
