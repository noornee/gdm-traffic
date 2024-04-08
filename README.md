# Description

**GDM-Traffic (GoogleDistanceMatrix-Traffic)**

GDM-Traffic is a Go program that fetches live traffic data from the Google Maps Platform Distance Matrix API based on the specified origin and destination locations.
This was written with the golang standard library It does not require any external dependncy besides `golang`

# Building from source
```bash
git clone https://github.com/noornee/gdm-traffic

cd gdm-traffic

go build -o gdm-traffic

```

# USAGE
Ensure that you have set the `MATRIX_API_KEY` environment variable with your Google Maps API key:
```bash
export MATRIX_API_KEY=your-api-key
```

### Flags
`--origin`: Specifies the longitude and latitude of the origin location. Values should be separated by a comma.

`--dest`: Specifies the longitude and latitude of the destination location. Values should be separated by a comma.

`--dit`:  Prints duration in traffic



### Run
```bash
./gdm-traffic --origin='123.4,-123.456' --dest='1234.012,-123.567'

# RESPONSE
[INFO]: Fetching data from google maps api...
[INFO]: {
   "destination_addresses" : 
   [
      "Lorem ipsum street"
   ],
   "origin_addresses" : 
   [
      "Lorem dolor street"
   ],
   "rows" : 
   [
      {
         "elements" : 
         [
            {
               "distance" : 
               {
                  "text" : "2.9 km",
                  "value" : 2889
               },
               "duration" : 
               {
                  "text" : "8 mins",
                  "value" : 466
               },
               "duration_in_traffic" : 
               {
                  "text" : "8 mins",
                  "value" : 459
               },
               "status" : "OK"
            }
         ]
      }
   ],
   "status" : "OK"
}
```

# Get the duration in traffic by specfying the --dit flag

```bash
./gdm-traffic --origin='123.4,-123.456' --dest='1234.012,-123.567' --dit

# RESPONSE
[INFO]: Fetching data from google maps api...
[INFO]: Travel time for origin 123.4,-123.456 and destination 1234.012,-123.567 is 8 mins

```



## Note ⚠
jsyk, the coords i passed to the flags are not valid, those are dummy coords so you'll need to pass valid ones
