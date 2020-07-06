package ioutil

import (
    "io/ioutil"

    "github.com/markbates/pkger"
)

func ReadFile(filename string) ([]byte, error) {
    f, err := pkger.Open(filename)
    if err != nil {
        return nil, err
    }

    b, err := ioutil.ReadAll(f)
    if err != nil {
        return nil, err
    }

    return b, nil
}
