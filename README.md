# log-collector version

## How to use
- Install the collector-prerequisite
```shell
helm install lsctest collector-prerequisite-1.0.0.tgz -f collector-prerequisite/values.yaml -n collect
```

## About tgzdownload
### This is a independent nodejs project for download log-collector build tgz package.
### How to use:	
1.  Make sure your node version is >=16.10.0
2.  Into log-collector/tgzdownload and run: 

#### `npm install`
#### `npm run dev`

3.  Open [http://localhost:3000](http://localhost:3000) in browser.
