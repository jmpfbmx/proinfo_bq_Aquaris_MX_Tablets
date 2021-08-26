WHAT IS THIS?
=============

Custom PROINFO binary partition for recover Serial Number losed meanwhile doing Format ALL at SP Flash Tool. Supported devices:
 * bq Aquaris M8
 * bq Aquaris M10 HD (Tested and working)
 * bq Aquaris M10 FHD
 * bq Aquaris M10 LTE

WHAT IS NEEDED?
===============

 * An hex-editor. (For example wxHexEditor)
 * TWRP for flashing the partition, since we will use dd.
 * Have serial number writen on a paper or the tablet with you.

HOW TO MODIFY PROINFO AND WRITE OUR SN AT IT?
=============================================

 - 1º Download <a href="https://raw.githubusercontent.com/jmpfbmx/proinfo_bq_Aquaris_MX_Tablets/master/PROINFO_BQ_MX.bin">PROINFO_BQ_MX.bin</a>

 - 2º Open your hex-editor and load PROINFO_BQ_MX.bin

 - 3º Modify the following values:<br>
   ![1st_IMAGE](https://i.imgur.com/5FPglKA.jpeg)

 - 4º In the case of our Serial Number is FA123456, we need to replace 30 values and set the following:
   * F == 46
   * A == 41
   * 1 == 31
   * 2 == 32
   * 3 == 33
   * 4 == 34
   * 5 == 35
   * 6 == 36<br>

    ![2nd_IMAGE](https://i.imgur.com/priDwnZ.png)

 - 5º Save the modifications done at PROINFO_BQ_MX.bin and copy it to your tablet.
 
 - 6º Reboot to TWRP and flash it:

   IMPORTANT NOTICE:<br>
     IF YOU COPIED TO SDCARD CHANGE PATH WITH /sdcard, IF YOU COPIED TO EXTERNAL SD CHANGE <PATH> WITH /external_sd

   * bq Aquaris M8:<br>
        ```
        adb shell
        dd if=/PATH/PROINFO_BQ_MX.bin of=/dev/block/platform/mtk-msdc.0/11230000.MSDC0/by-name/proinfo
        ```

   * bq Aquaris M10 HD:<br>
      OFFICIAL TWRP (LP AND MM FW):<br>
        ```
        adb shell
        dd if=/PATH/PROINFO_BQ_MX.bin of=/dev/block/platform/mtk-msdc.0/by-name/proinfo
        ```

      UNOFFICIAL TWRP (OREO FW):<br>
        ```
        adb shell
        dd if=/PATH/PROINFO_BQ_MX.bin of=/dev/block/platform/mtk-msdc.0/11230000.MSDC0/by-name/proinfo
        ```

   * bq Aquaris M10 FHD:<br>
        ```
        adb shell
        dd if=/PATH/PROINFO_BQ_MX.bin of=/dev/block/platform/mtk-msdc.0/by-name/proinfo
        ```

   * bq Aquaris M10 4G:<br>
        ```
        adb shell
        dd if=/PATH/PROINFO_BQ_MX.bin of=/dev/block/platform/mtk-msdc.0/11230000.msdc0/by-name/proinfo
        ```

## Author:
 - JMPFBMX
 - coyuya

### Help:
 * If you need help, if you don't know which is the value of some letter or if you don't know how to do it, open a github issue.
