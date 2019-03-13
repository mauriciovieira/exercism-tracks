class Luhn
  def self.valid? number
    digits = number.gsub(/\s/, '')
    if digits.size < 2 || digits.match(/^\d+$/).nil?
      return false
    end

    sum = 0

    digits = number.scan(/\d/).map(&:to_i)
    digits.reverse.each_with_index do |digit, i|
      if i.odd?
        sum += if digit > 4 then digit * 2 - 9 else digit * 2 end
      else
        sum += digit
      end
    end

    sum % 10 == 0
  end
end