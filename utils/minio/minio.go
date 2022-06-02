/**
 * @Time    :2022/5/23 9:03
 * @Author  :Xiaoyu.Zhang
 */

package minio

import (
	"context"
	"errors"
	"github.com/melf-xyzh/gin-start/global"
	"github.com/minio/minio-go/v7"
	"log"
	"sync"
)

var (
	ctx  = context.Background()
	once sync.Once
)

// initBucket
/**
 *  @Description: 初始化存储桶
 *  @param bucketName
 *  @param location
 *  @return err
 */
func initBucket(bucketName, location string) (err error) {
	if location == "" {
		location = "us-east-8"
	}
	exists, errBucketExists := global.MioIO.BucketExists(ctx, bucketName)
	if errBucketExists == nil && exists {
		log.Printf("存储桶%s已存在", bucketName)
	} else {
		err = global.MioIO.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			err = errors.New("创建存储桶" + bucketName + "失败")
		} else {
			log.Printf("创建存储桶%s成功", bucketName)
		}
	}
	return
}

// UpLoad
/**
 *  @Description: 上传
 *  @param objectName aa/bb/test.sql
 *  @param filePath D:\\test.sql
 *  @param contentType application/zip
 */
func UpLoad(objectName, filePath, contentType string) (err error) {
	bucketName := global.V.GetString("MinIO.bucketName")
	once.Do(func() {
		err = initBucket(bucketName, "")
		if err != nil {
			panic(err)
		}
	})
	_, err = global.MioIO.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	return
}

// DownLoad
/**
 *  @Description: 下载
 *  @param objectName
 *  @param filePath
 *  @param contentType
 */
func DownLoad(objectName, filePath string) (err error) {
	bucketName := global.V.GetString("MinIO.bucketName")
	err = global.MioIO.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	return
}

// Delete
/**
 *  @Description: 删除文件
 *  @param objectName
 *  @return err
 */
func Delete(objectName string) (err error) {
	bucketName := global.V.GetString("MinIO.bucketName")
	//删除一个文件
	err = global.MioIO.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{GovernanceBypass: true})
	return
}

// Exist
/**
 *  @Description: 文件是否存在
 *  @param objectName
 *  @param filePath
 *  @param contentType
 */
func Exist(objectName string) (exist bool) {
	bucketName := global.V.GetString("MinIO.bucketName")
	_, err := global.MioIO.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return false
	}
	return true
}
