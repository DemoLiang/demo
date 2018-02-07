#!/usr/bin/env python
# -*- coding:utf-8 -*-

import urllib2
import time
import re
from bs4 import BeautifulSoup


class HtmlDownloader(object):
    header = {'Cookie': 'AD_RS_COOKIE=20083363',
              'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) \
              AppleWeb\Kit/537.36 (KHTML, like Gecko)\
              Chrome/58.0.3029.110 Safari/537.36'}
    def download(self,url):
        if url is None:
            raise Exception('url is None')
        # print url
        request = urllib2.Request(url,None,HtmlDownloader.header)
        try:
            resp = urllib2.urlopen(request)
            # print resp.getcode()
            if resp.getcode()!= 200:
                time.sleep(5)
                return self.download(url)
            else:
                return resp.read()
        except urllib2.URLError,e:
            print e
            time.sleep(5)
            return self.download(url)
    def readhtml(self,filename):
        file_object = open(filename)
        try:
            all_the_text = file_object.read()
        finally:
            file_object.close()
        return all_the_text



class HtmlParser(object):
    def has_tag(self,tag):
        return tag.has_attr('span')
    def region_parser(self,html_content):
        if html_content is None:
            raise Exception('html is None')
        soup = BeautifulSoup(html_content,'html.parser')
        for tag in soup.find_all(class_="MsoNormal"):
            # print tag.get_text()
            id = tag.get_text().split(" ")[0].strip()
            name = tag.get_text().split(" ")[1].strip()
            print id+"-->"+name
    def contry_parse(self,html_content):
        if html_content is None:
            raise Exception('html is None')
        soup = BeautifulSoup(html_content, 'html.parser')
        for tag in soup.find_all(class_="wikitable sortable"):
            # print tag
            # print tag.select('td')
            i = 0
            for td in tag.select('td'):
                if i % 5 == 0:
                    print "id-->"+td.get_text().strip()
                elif i % 5 == 4:
                    print "name-->"+td.get_text().strip()
                # print str(i)+"----"+td.get_text().strip()
                i = i + 1
            # code = tag.get_text().split(" ")[0].strip()
            # name = tag.get_text().split(" ")[3].strip()
            # print code + "-->" + name
    def contry_ipaddrlink_parse(self,html_content):
        if html_content is None:
            raise Exception('html is None')
        soup = BeautifulSoup(html_content, 'html.parser')
        for tag in soup.find_all(href=re.compile(u'http://ipblock.chacuo.net/view/.*')):
            print  tag.get_text()+"-->"+tag.get('href')
            # html_content = html_downloader.download(tag.get('href'))
            # print tag

    def ipaddress_parse(self,html_content):
        if html_content is None:
            raise Exception('html is None')
        soup = BeautifulSoup(html_content, 'html.parser')
        for tag in soup.find_all('pre'):
            # print tag.get_text()+"-->"
            return tag.get_text()

    def ipaddress_parse_text(self,html_content):
        if html_content is None:
            raise Exception('html is None')
        soup = BeautifulSoup(html_content, 'html.parser')
        for tag in soup.find_all(href=re.compile(u'http://ipblock.chacuo.net/view/.*')):
            # print re.sub(r'view/c_', "down/t_txt=c_", tag.get('href'))
            content = html_downloader.download(re.sub(r'view/c_', "down/t_txt=c_", tag.get('href')))
            try:
                print tag.get_text()
                contentstr = self.ipaddress_parse(content)
                # print contentstr
                for ipdata in contentstr.split('\r\n'):
                    data = ipdata.split('\t')
                    # print data[1]
                    # print data[0]
                    # print '-->'+ipdata
                    # print ipdata.split('\t')[0]
                    # print data[0].strip()
                    if len(data) >3:
                        print '--->ip:'+data[0]+'--->mask:'+data[1]+'-->mask/len:'+data[2]+'-->num:'+data[3]
            except Exception,e:
                print "no data"

    def s_ipaddress_parse_text(self, html_content):
        if html_content is None:
            raise Exception('html is None')
        soup = BeautifulSoup(html_content, 'html.parser')
        for tag in soup.find_all(href=re.compile(u'http://ips.chacuo.net/view/.*')):
            # print re.sub(r'view/c_', "down/t_txt=c_", tag.get('href'))
            content = html_downloader.download(re.sub(r'view/s_', "down/t_txt=p_", tag.get('href')))
            # print tag.get('href')+"--->"
            try:
                print tag.get_text()
                contentstr = self.ipaddress_parse(content)
                # print contentstr
                for ipdata in contentstr.split('\r\n'):
                    # print ipdata
                    data = ipdata.split('\t')
                    # print data
                #     # print data[1]
                #     # print data[0]
                #     # print '-->'+ipdata
                #     # print ipdata.split('\t')[0]
                #     # print data[0].strip()
                    if len(data) > 2:
                        print '--->ip:' + data[0] + '--->mask:' + data[1] + '-->num:' + data[2]
            except Exception, e:
                print "no data"
    def isp_ipaddress_parse_text(self, html_content):
        if html_content is None:
            raise Exception('html is None')
        soup = BeautifulSoup(html_content, 'html.parser')
        for tag in soup.find_all(href=re.compile(u'http://ipcn.chacuo.net/view/.*')):
            # print re.sub(r'view/c_', "down/t_txt=c_", tag.get('href'))
            content = html_downloader.download(re.sub(r'view/i_', "down/t_txt=c_", tag.get('href')))
            # print tag.get('href')+"--->"
            try:
                print tag.get_text()
                contentstr = self.ipaddress_parse(content)
                # print contentstr
                for ipdata in contentstr.split('\r\n'):
                    # print ipdata
                    data = ipdata.split('\t')
                    # print data
                #     # print data[1]
                #     # print data[0]
                #     # print '-->'+ipdata
                #     # print ipdata.split('\t')[0]
                #     # print data[0].strip()
                    if len(data) > 2:
                        print '--->ip:' + data[0] + '--->mask:' + data[1] + '-->num:' + data[2]
            except Exception, e:
                print "no data"

if __name__ == '__main__':
    html_downloader = HtmlDownloader()

    #region
    # html_content = html_downloader.download('http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201703/t20170310_1471429.html')
    # html_parser = HtmlParser()
    # html_parser.region_parser(html_content)

    #contry
    # html_content = html_downloader.readhtml('ISO3166-1.html')
    # html_parser = HtmlParser()
    # html_parser.contry_parse(html_content)

    #contry ip address link parse
    # html_content = html_downloader.download('http://ipblock.chacuo.net')
    # html_parser = HtmlParser()
    # html_parser.contry_ipaddrlink_parse(html_content)

    #contry ip address parse
    # html_content = html_downloader.download('http://ipblock.chacuo.net/down/t_txt=c_AO')
    # html_parser = HtmlParser()
    # html_parser.ipaddress_parse(html_content)

    # contry ip address parse to text
    # html_content = html_downloader.download('http://ipblock.chacuo.net')
    # html_parser = HtmlParser()
    # html_parser.ipaddress_parse_text(html_content)

    #cn s ipaddress parse
    html_content = html_downloader.download('http://ips.chacuo.net/')
    html_parser = HtmlParser()
    html_parser.s_ipaddress_parse_text(html_content)

    # cn s ipaddress parse
    html_content = html_downloader.download('http://ipcn.chacuo.net/')
    html_parser = HtmlParser()
    html_parser.isp_ipaddress_parse_text(html_content)

    # print html_content
