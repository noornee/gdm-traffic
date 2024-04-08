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

`--tt`:  Prints duration in traffic/travel time

### Run
```bash
./gdm-traffic --origin='123.4,-123.456' --dest='1234.012,-123.567'
```

### RESPONSE
```
# would update this in a few minutes
```

