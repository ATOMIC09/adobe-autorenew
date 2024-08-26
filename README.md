# ⚠️ Update
ทางมหาวิทยาลัยได้มีการเปลี่ยนเงื่อนไขจากการต่ออายุใบอนุญาตจาก**รายสัปดาห์**เป็น**รายเดือน**แล้วแต่โค้ดไม่ได้มีการอัปเดต ดังนั้น**ไม่ควรใช้อย่างยิ่ง**

The university has changed the license renewal terms from **weekly** to **monthly**, but the code hasn't been updated yet. So, **it’s strongly advised not to use it.**

# 🚀 How to configure and run the script
*(This script has been tested and run on-premises)*

1. Install **Docker**

2. Create `.env` file
```
MY_USERNAME=your_username
MY_PASSWORD=your_password
```

3. Clone this repository or create your `Dockerfile` and use `main.go` file from [this original source](https://gist.github.com/buratud/ed6e786287c3a42ef70da2f85c311244)

4. Build the image by using this command
`docker build -t adobe-autorenew --build-arg ENV_FILE=.env .`

5. Run the image!
`docker run -d --restart=unless-stopped --env-file .env -e TZ=Asia/Bangkok adobe-autorenew`

<hr/>
 
`-d` is "detached mode." When you run a container in detached mode, it means the container runs in the background and doesn't block the terminal. 

`--restart=unless-stopped` always restarts the container unless the user explicitly stops it.
