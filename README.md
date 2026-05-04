What Is Lingo?

As Of Now Lingo Is A Set Of Functions or Extensions That Can Help You Querying Data from any Structures. Something Like Linq In C# Or Streams In Java. Detailed Documentation Available In Wiki Section. This Project Is An Open Source Project Under No Guarantees, Though We Use Test Driven Approach To Ensure Stable Releases.



Rich Syntax

Ease Of Use

Open Source

Fast Development

Integrated Set Of Tools


Our Benchmark Shows The Following Lines Of Code Took About 8 seconds To Run In A Slice Of 50,000,000 records. And Results Were As Expected. It Sounds Reasonable.
	
	Benchmark Pc Specs:
	Laptop 
	Intel Core I7 12700
	16Gb Of Ram
	SSD
	Linux Ubuntu
    8 Seconds
The Benchmark Included In The Test File


	res, err := AllOrDefault(From(items).Where("Flag", true).Filter(func(item ComplexObjectToSearch) bool {
		return item.Id > 200000
	}))

	if err != nil {
		b.Error(err)

	}

	if Any(*res, func(item ComplexObjectToSearch) bool {
		return !item.Flag
	}) {
		b.Error("Wrong Data Fetched")
	}
