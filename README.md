WHAT IS THIS?
=============

Custom PROINFO binary partition for recover Serial Number losed meanwhile doing Format ALL at SP Flash Tool. Supported devices:
 * bq Aquaris M8
 * bq Aquaris M10 HD (Tested and working)
 * bq Aquaris M10 FHD
 * bq Aquaris M10 LTE


WHAT IS write_sn?
=================
 * Is a custom binary which source have been written in golang by @coyuya, which modifies PROINFO_BQ_MX.bin writing your serial number in it.<br>


WHAT IS NEEDED?
===============

 * A computer with Windows or Linux.
 * A CMD or Terminal.
 * TWRP for flashing the partition, since we will use dd.
 * Have serial number writen on a paper or the tablet with you.

HOW TO USE WRITE_SN?
===========================================================================

 <h4>LINUX:</h4>

 - 1º Clone proinfo_bq_Aquaris_MX_Tablets or download it from github and extract it.
    ```
    git clone -b go https://github.com/jmpfbmx/proinfo_bq_Aquaris_MX_Tablets && cd proinfo_bq_Aquaris_MX_Tablets
    ```

 - 2º Give correct permissions to write_sn binary:
    ```
    chmod 777 write_sn
    ```

 - 3º Execute it:<br>
    ```
    ./write_sn
    ```
   ![1st_IMAGE](https://i.imgur.com/y5zstit.png)

 - 4º Write our Serial Number (For example: FA123456):<br>

    ![2nd_IMAGE](https://i.imgur.com/XOqI7wJ.png)

 - 5º Copy PROINFO_BQ_MX.bin to your tablet.
 
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

 <h4>WINDOWS:</h4>

 - 1º Clone proinfo_bq_Aquaris_MX_Tablets or download it from github and extract it.

 - 2º Open a CMD at proinfo_bq_Aquaris_MX_Tablets folder

 - 3º Execute write_sn:<br>
    ```
    .\write_sn.exe
    ```
   ![1st_IMAGE](https://i.imgur.com/HmIEzWY.png)

 - 4º Write our Serial Number (For example: FA123456):<br>

    ![2nd_IMAGE](https://i.imgur.com/K7EW8rq.png)

 - 5º Copy PROINFO_BQ_MX.bin to your tablet.
 
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

## Authors:
 - JMPFBMX
 - coyuya
