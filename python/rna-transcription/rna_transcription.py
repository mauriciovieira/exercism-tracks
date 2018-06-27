def to_rna(dna_strand):
    transcription = {'A': 'U', 'T': 'A', 'G': 'C', 'C': 'G'}
    return ''.join([transcription[nucleotide] for nucleotide in dna_strand])
