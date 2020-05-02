class Isogram
  def self.isogram?(word)
    freq = Hash.new(0)
    word.downcase.scan(/\w/).each do |letter|
      freq[letter] += 1
    end
    freq.values.all? { |count| count == 1 }
  end
end