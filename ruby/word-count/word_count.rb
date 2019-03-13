class Phrase
  def initialize(phrase)
    @phrase = phrase
  end

  def word_count
    counts = Hash.new(0)
    @phrase.downcase.scan(/\w+\'t|\w+/).each do |word|
      counts[word] += 1
    end
    counts
  end
end