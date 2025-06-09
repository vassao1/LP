potencia :: Int -> Int -> Int
potencia x y = x^y

main :: IO ()
main = do
    putStrLn "Insira a base:"
    base <- getLine
    putStrLn "Insira o expoente:"
    expoente <- getLine
    let baseInt = read base :: Int
    let expoenteInt = read expoente :: Int
    let resultado = potencia baseInt expoenteInt
    putStrLn ("O resultado de " ++ show baseInt ++ " elevado a " ++ show expoenteInt ++ " eh " ++ show resultado)