class Raindrops
  def self.convert(number)
    answers = []

    answers << "Pling" if (number % 3).zero?
    answers << "Plang" if (number % 5).zero?
    answers << "Plong" if (number % 7).zero?

    if answers.empty?
      number.to_s
    else
      answers.join('')
    end
  end
end