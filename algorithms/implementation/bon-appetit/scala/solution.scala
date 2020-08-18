import java.io._
import java.math._
import java.security._
import java.text._
import java.util._
import java.util.concurrent._
import java.util.function._
import java.util.regex._
import java.util.stream._
import scala.collection.immutable._
import scala.collection.mutable._
import scala.collection.concurrent._
import scala.concurrent._
import scala.io._
import scala.math._
import scala.sys._
import scala.util.matching._
import scala.reflect._

object Solution {

    // Complete the bonAppetit function below.
    def bonAppetit(bill: Array[Int], k: Int, b: Int) {

        val equalPay: Int = (bill.foldLeft(0)(_+_) - bill(k)) / 2
        val diff: Int = equalPay - b

        diff match {
            case 0 => println("Bon Appetit")
            case _ => println(diff.abs)
        }
    }

    def main(args: Array[String]) {
        val nk = StdIn.readLine.replaceAll("\\s+$", "").split(" ")

        val n = nk(0).toInt

        val k = nk(1).toInt

        val bill = StdIn.readLine.replaceAll("\\s+$", "").split(" ").map(_.trim.toInt)
        val b = StdIn.readLine.trim.toInt

        bonAppetit(bill, k, b)
    }
}

