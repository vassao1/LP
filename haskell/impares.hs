impares :: Int -> [Int]
impares n = take n [x | x <- [1..], odd x]
  
main :: IO ()
main = do
  putStrLn "Digite um numero:"
  input <- getLine
  let num = read input :: Int
  putStrLn $ "Os " ++ show num ++ " primeiros naturais impares sao: " ++ show (impares num)