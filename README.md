# colormap2label 

This tools converts colormap images into grayscale label images in parallel.
It is useful to create datasets for semantic segmentation.


## Installation

### Requirements

-   Golang for building 

### Steps

										    
1.  Compile tools if you need 
											    
```shell
$ go build -o colomap2label
```


## Usage

### Convet BDD100K laneline

```shell
$ ./colomap2label -csv csv/BDD100k_lane.csv -i {Input Directory} -o {Output Directory}
```
