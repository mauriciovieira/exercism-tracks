class ArmstrongNumbers
  def self.include?(number)
    digits = number.digits
    size = digits.size
    digits.sum{|digit| digit ** size} == number
  end
end