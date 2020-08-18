import java.io._
import java.math._
import java.security._
import java.text._
// import java.util._
import java.util.concurrent._
import java.util.function._
import java.util.regex._
import java.util.stream._
import scala.collection.immutable._
// import scala.collection.mutable._
// import scala.collection.concurrent._
import scala.concurrent._
import scala.io._
import scala.math._
import scala.sys._
import scala.util.matching._
import scala.reflect._

object Solution {

    // Complete the sockMerchant function below.
    def sockMerchant(n: Int, ar: Array[Int]): Int = {
        val hmap: Map[Int, Int] = ar.groupBy(identity).map{ case (x, y) => x -> y.size }
        hmap.values.foldLeft(0)(_ + _/2)
    }

    def main(args: Array[String]) {
        val stdin = scala.io.StdIn

        val printWriter = new PrintWriter(sys.env("OUTPUT_PATH"))

        val n = stdin.readLine.trim.toInt

        val ar = stdin.readLine.split(" ").map(_.trim.toInt)
        val result = sockMerchant(n, ar)

        printWriter.println(result)

        printWriter.close()
    }
}

