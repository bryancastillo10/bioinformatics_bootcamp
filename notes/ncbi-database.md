# NCBI Gene Search, Accession Number, BLAST, and Data Download

## 1. NCBI Gene Search

NCBI is a major biological database resource used to search genes, nucleotide records, protein sequence, related literatures, and other biological information.

- Genes can be searched via several keywords (eg. gene name ,organism)
- Search results may include multiple related records
- To retrieve the desired gene, it is important to confirm every details that matches what you need

## 2. Accession Number

An accession number is a unique identifier assigned to a biological sequence or database record

### Importance

- helps identifying an exact record
- avoids confusion to genes that have similar names
- can be used directly for sequence retrieval or BLAST searches

## 3. BLAST (Basic Local Alignment Search Tool)

BLAST is a fundamental bioinformatics algorithm used to compare a primary biological sequence (query) against a database of sequences to identify regions of similarity. It rapidly finds homologous nucleotide or protein sequences, helping to determine evolutionary relationships and identify unknown sequences.

There are several types of BLAST such as:

1. **BLASTn** compares nucleotide (DNA sequences)
2. **BLASTp** compares protein sequences
3. **BLASTx** translates DNA query to protein and compare protein databases
4. **tBLASTn** compares protein query to translated nucleotide database

## 4. Downloading Raw Data

Sequence data can be downloaded from NCBI for later analysis.

### Common download formats

1. FASTA
2. GenBank
3. sometimes other metadata-rich formats

## Key Terms

- NCBI
- gene record
- accession number
- nucleotide
- protein
- BLAST
- FASTA
- GenBank


### Practical
1. Explore NCBI Database:
- Go to the NCBI website (https://www.ncbi.nlm.nih.gov) and navigate to the "All Databases" tab
- Explore the different databases available such as PubMed, GenBank, and BLAST
- Click on each database to learn more about its content and purpose

2. Retrieve an accession number from GenBank:
- Go to the GenBank database on the NCBI website
- Use the search bar to find a specific gene or organism you are interested in
- Click on a result to view details of the sequence
- Copt the accession number of the sequence

3. Use BLAST to find similar sequences:
- Go to the BLAST (Basic Local Alignment Search Tool) page on the NCBI website
- Choose the type of BLAST you want to perform (eg. BLASTn for nucleotide sequences)
- Enter the sequence you retrieved from GenBank or a sequence of your own
- Run the BLAST search and explore the results to find similar sequences

4. Download a Sequence from GenBank
- Go back to the GenBank database and enter the accession number of the sequence found earlier
- Locate the "Send to" option and chose the "File" to download the sequence in a  format of choice (eg. FASTA)

Practical Search: bzip2 which has accession number of **AT2G18160**