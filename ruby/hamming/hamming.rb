class Hamming
  def self.compute(a, b)
    a_arr = a.split('')
    b_arr = b.split('')
    raise ArgumentError if a.size != b.size
    a_arr.zip(b_arr).count{|pair| pair[0] != pair[1]}
  end
end