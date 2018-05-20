class Complement
  @@complements = {
    "G" => "C",
    "C" => "G",
    "A" => "U",
    "T" => "A"
  }

  def self.of_dna(dna_sequence)
    return "" unless dna_sequence.chars.uniq.reject {|l| @@complements.keys.include? l }.empty?
    dna_sequence.chars.map {|l| @@complements[l] }.join ''
  end
end

module BookKeeping
  VERSION = 4
end