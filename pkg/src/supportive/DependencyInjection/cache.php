<?php

declare (strict_types=1);
namespace DEPTRAC_INTERNAL;

use Qossmic\Deptrac\Core\Ast\Parser\Cache\AstFileReferenceCacheInterface;
use Qossmic\Deptrac\Core\Ast\Parser\Cache\AstFileReferenceDeferredCacheInterface;
use Qossmic\Deptrac\Core\Ast\Parser\Cache\AstFileReferenceFileCache;
use Qossmic\Deptrac\Core\Ast\Parser\Cache\CacheableFileSubscriber;
use Qossmic\Deptrac\Supportive\Console\Application;
use DEPTRAC_INTERNAL\Symfony\Component\DependencyInjection\Loader\Configurator\ContainerConfigurator;
use function DEPTRAC_INTERNAL\Symfony\Component\DependencyInjection\Loader\Configurator\service;
return static function (ContainerConfigurator $container) : void {
    $services = $container->services();
    $services->defaults()->public();
    $services->set(AstFileReferenceFileCache::class)->args(['%cache_file%', Application::VERSION]);
    $services->alias(AstFileReferenceDeferredCacheInterface::class, AstFileReferenceFileCache::class);
    $services->alias(AstFileReferenceCacheInterface::class, AstFileReferenceDeferredCacheInterface::class);
    $services->set(CacheableFileSubscriber::class)->args([service(AstFileReferenceFileCache::class)])->tag('kernel.event_subscriber');
};
