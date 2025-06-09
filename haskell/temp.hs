celsius :: Float -> Float
celsius n = ((n-32)*5)/9
  
main :: IO ()
main = do
  putStrLn "Digite um número:"
  input <- getLine
  let fahrenheit = read input :: Float
  putStrLn $ "A temperatura " ++ show fahrenheit ++ " em Celsius eh: " ++ show (celsius fahrenheit)