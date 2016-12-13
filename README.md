##Rp Converter
convert float64 into string (in Bahasa)
```
if x <= 0{ return "NOL RUPIAH" }
```
only supports until x < 1 trillion

Usage example:
```
rp := rpconv.Convert(993111550123)
fmt.Println(rp)
```
This should print 
```
SEMBILAN RATUS SEMBILAN PULUH TIGA MILYAR SERATUS SEBELAS JUTA LIMA RATUS LIMA PULUH RIBU SERATUS DUA PULUH TIGA RUPIAH
```

