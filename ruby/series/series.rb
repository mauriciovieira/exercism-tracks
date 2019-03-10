class Series
  def initialize(number)
    @digits = number.split ''
  end

  def slices(n)
    raise ArgumentError if n > @digits.size

    slices = []
    @digits.each_with_index do |e, i|
      slices << @digits.slice(i, n).join if i+n <= @digits.size
    end
    slices
  end
end
