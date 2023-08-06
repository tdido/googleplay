# google play services 9

## Android Studio

starting with this:

<http://dl.google.com/android/repository/sys-img/google_apis_playstore/x86-24_r12.zip>

this Go program fails:

~~~go
package main

import (
   "fmt"
   "net/http"
   "time"
)

func main() {
   req, err := http.NewRequest("HEAD", "http://dl.google.com", nil)
   if err != nil {
      panic(err)
   }
   for r := 0; r <= 64; r++ {
      req.URL.Path = "/android/repository/sys-img/google_apis_playstore"
      req.URL.Path += fmt.Sprintf("/x86-23_r%02v.zip", r)
      res, err := http.DefaultClient.Do(req)
      if err != nil {
         panic(err)
      }
      fmt.Println(res.Status, req.URL)
      if res.StatusCode == http.StatusOK {
         break
      }
      time.Sleep(time.Second)
   }
}
~~~

## Open GApps

https://sourceforge.net/projects/opengapps/files/x86

~~~
Core\gmscore-x86.tar.lz
~~~

then:

~~~
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk
~~~

result:

~~~
versionCode='18382005'
versionName='18.3.82 (000700-260264002)'
~~~

<https://sourceforge.net/projects/opengapps/files/x86_64>

~~~
versionCode='19831014'
versionName='19.8.31 (020800-284611645)'
~~~

https://github.com/opengapps/opengapps/issues/982

## FlameGApps

~~~
name='com.google.android.gms'
versionCode='19829028'
versionName='19.8.29 (100400-282600551)'
~~~

https://github.com/flamegapps/flamegapps/issues/11

## BiTGApps

~~~
name='com.google.android.gms'
versionCode='17455017'
versionName='17.4.55 (040306-248795830)'
~~~

https://bitgapps.io

## MindTheGapps

~~~
name='com.google.android.gms'
versionCode='11976970'
versionName='11.9.76 (970-184349000)'
~~~

https://mindthegapps.com

## NikGApps

~~~
name='com.google.android.gms'
versionCode='19528039'
versionName='19.5.28 (100408-273329419)'
~~~

https://nikgapps.com