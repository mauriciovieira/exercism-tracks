class Series
  def initialize(number)
    @digits = number.chars
  end

  def slices(n)
    raise ArgumentError if n > @digits.size
    @digits.each_cons(n).map(&:join)
  end
end
