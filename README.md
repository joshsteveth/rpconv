##Rp Converter
convert float64 into string (in Bahasa)
only supports until 0 <= x < 10^12

Usage example:
```
if rp, err := rpconv.Convert(993111550123); err != nil{
	fmt.Println(rp)
}
```
This should print 
```
SEMBILAN RATUS SEMBILAN PULUH TIGA MILYAR SERATUS SEBELAS JUTA LIMA RATUS LIMA PULUH RIBU SERATUS DUA PULUH TIGA RUPIAH
```

