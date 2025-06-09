primo :: Int -> Bool
primo n
  | n < 2     = False
  | n == 2    = True
  | otherwise = null [x | x <- [2..(n-1)], n `mod` x == 0]

main :: IO ()
main = do
  putStrLn "Digite um numero:"
  input <- getLine
  let numero = read input :: Int
  if primo numero
    then putStrLn $ show numero ++ " eh primo."
    else putStrLn $ show numero ++ " não eh primo."