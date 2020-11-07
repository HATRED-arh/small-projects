import java.time.temporal.ChronoUnit
import java.time.LocalDate
import java.io.File
import scala.io.Source.fromFile
import java.io._
import java.time.format.DateTimeParseException

object Count extends App {
  val filename = "date.dat"
  if (!new File(filename).exists) {
    println("File created")
    defaultFileCreator(filename)
    System.exit(0)
  }

  dateChecker(readFile(filename), filename)
  val startDate = LocalDate.parse(readFile(filename))
  val currentDate = LocalDate.now
  var days = ChronoUnit.DAYS.between(startDate, currentDate)
  println(s"Days passed: $days")
  days.toDouble
  days = days / 4 + 8
  println(s"This many: $days")

  def dateChecker(input: String, name: String): Unit = {
    try {
      LocalDate.parse(input)
    } catch {
      case _: DateTimeParseException =>
        println(
          "Data structure in file is wrong\nErasing data\nWriting current date to file"
        )
        defaultFileCreator(name)
        System.exit(0)
    }
  }
  def readFile(name: String): String = {
    val red = fromFile(name)
    val str = red.mkString
    red.close()
    str
  }
  def defaultFileCreator(name: String): Unit = {
    val data = LocalDate.now().toString
    val buf = new BufferedWriter(new FileWriter(name))
    buf.write(data)
    buf.close()
  }
}
